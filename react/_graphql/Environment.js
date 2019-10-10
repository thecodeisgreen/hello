const R = require('ramda');
import Tokens from '../lib/Tokens';
import assert from 'assert';
import { execute } from 'apollo-link';
import { SubscriptionClient } from 'subscriptions-transport-ws';
import { WebSocketLink } from 'apollo-link-ws';

const {
  Environment,
  Network,
  RecordSource,
  Store,
} = require('relay-runtime');

const GQLEnvironment = async () => {
  const sendRequest = async (url, headers, body, refresh_token = false) => {
    let access_token = await Tokens.getAccessToken(refresh_token);
    const response = await fetch(url, {
      method: 'POST',
      credentials: 'same-origin',
      headers: {
        ...headers,
        //'Authorization': `Bearer ${access_token}`
      },
      body
    });

    if (response.status === 200) {
      return await response.json();
    }

    if (response.status === 401) {
      assert(!refresh_token, 'can not refresh access token');
      return await sendRequest(url, headers, body, true);
    }

    return new Error('graphql request failed');
  };

  const fetchQuery = async (operation, variables, cacheConfig) => {
    let body = new FormData();
    body = JSON.stringify({
      query: operation.text,
      variables,
    });
    let headers = {
      Accept: '*/*',
      'Content-Type': 'application/json'
    };

    return await sendRequest('/graphql', headers, body);
  };

  const graphqlWsUrl = async () => {
    try {
    const response = await fetch('/graphqlwsurl');
    return R.prop('url', await response.json());
    } catch(err) {
      return "ws://localhost:8080"
    }
  };

  const subscriptionClient = new SubscriptionClient(await graphqlWsUrl(), {
    reconnect: true,
  });

  const subscriptionLink = new WebSocketLink(subscriptionClient);

  // Prepar network layer from apollo-link for graphql subscriptions
  const networkSubscriptions = (operation, variables) =>
    execute(subscriptionLink, {
      query: operation.text,
      variables,
    });
  
  return new Environment({
    network: Network.create(fetchQuery, networkSubscriptions),
    store: new Store(new RecordSource()),
  });
};

export default GQLEnvironment;