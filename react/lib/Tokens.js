const R = require('ramda');
import fetch from 'isomorphic-fetch';
import $url from 'url';

const Tokens = (() => {
  var vars = {
    clientId: null,
    key: null,
    tokens: {},
    useToken: true
  };

  const buildUrl = (pathname, query) => {
    return $url.format({ pathname, query });
  };

  const setKey = (client_id, client_secret) => {
    vars = R.compose(
      R.assoc('client_id', client_id),
      R.assoc('client_secret', client_secret)
    )(vars);
  };

  const getKey = (k) => R.prop(k)(vars)

  var getAccessToken = async (refresh = false) => {
    if (refresh) return getAccessTokenRefresh();

    if (R.prop('useToken', vars) === false) {
      return Promise.resolve('');
    }

    if (R.compose(R.not, R.isEmpty, R.prop('tokens'))(vars)) return Promise.resolve(R.path(['tokens', 'access_token'], vars));

    const formData = new FormData();
    formData.append('grant_type', 'client_credentials')
    formData.append('client_id', getKey('client_id'));
    formData.append('client_secret', getKey('client_secret'));
    formData.append('scope', 'react-app');

    try {
      const response = await fetch(
        buildUrl('/o/token'),
        {
          method: 'POST',
          body: formData
        }
      );

      vars = R.assoc('tokens', await response.json(), vars);
      return R.path(['tokens', 'access_token'], vars);
    } catch(err) {
      console.log(err.stack)
    }
   };

  const getAccessTokenRefresh = async () => {
    const formData = new FormData();
    formData.append('grant_type', 'refresh_token');
    formData.append('refresh_token', R.path(['tokens', 'refresh_token'], vars));
    formData.append('client_id', getKey('client_id'));
    formData.append('client_secret', getKey('client_secret'));

    const response = await fetch(
      buildUrl('/o/token'),
      {
        method: 'POST',
        body: formData
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