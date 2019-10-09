const R = require('ramda');

import React, { useState, useEffect, useRef } from 'react';

import { graphql } from 'react-relay';


import {
  Query
} from '../_graphql';

const UserQuery = graphql`
  query UserQuery {
    user {
      id
      email
    }
  } 
`;

const User = () => (
  <Query
    query={UserQuery}
  >
    {({user}) => (
      <div>{user.email}</div>
    )}
  </Query>
)

export default User