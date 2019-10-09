const R = require('ramda');
import fetch from 'isomorphic-fetch';
import Url from 'url';

const Tokens = (() => {
  var vars = {
    clientId: null,
    key: null,
    tokens: {},
    useToken: true
  };

  const buildUrl = (pathname, query) => {
    return Url.format({ pathname, query });
  };

  const setKey = (clientId, key) => {
    vars = R.compose(
      R.assoc('key', key),
      R.assoc('clientId', clientId)
    )(vars);
  };

  var getAccessToken = async (refresh = false) => {
    if (refresh) return getAccessTokenRefresh();

    if (R.prop('useToken', vars) === false) {
      return Promise.resolve('');
    }

    if (R.compose(R.not, R.isEmpty, R.prop('tokens'))(vars)) return Promise.resolve(R.path(['tokens', 'access_token'], vars));

    const response = await fetch(
      buildUrl('/u/token'),
      {
        method: 'POST',
        body: JSON.stringify({ 
          grant_type: 'client_credentials'
        }),
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Basic ' + btoa(`${R.prop('clientId', vars)}:${R.prop('key', vars)}`),
          'Content-Type': 'application/json'
        }
      }
    );

    vars = R.assoc('tokens', await response.json(), vars);
    return R.path(['tokens', 'access_token'], vars);
  };

  const getAccessTokenRefresh = async () => {
    const response = await fetch(
      buildUrl('/u/token'),
      {
        method: 'POST',
        body: JSON.stringify({ 
          grant_type: 'refresh_token',
          refresh_token: R.path(['tokens', 'refresh_token'], vars)
        }),
        headers: {
          'Accept': 'application/json',
          'Authorization': 'Basic ' + btoa(`${R.prop('clientId', vars)}:${R.prop('key', vars)}`),
          'Content-Type': 'application/json'
        }
      }
    );
    const tokens = await response.json();
    vars = R.assoc('tokens', tokens, vars);

    return R.path(['tokens', 'access_token'], vars);
  };

  const setUseToken = (useToken) => {
    vars = R.assoc('useToken', useToken, vars);
  };

  return {
    getAccessToken,
    setKey,
    setUseToken
  };
})();

export default Tokens;