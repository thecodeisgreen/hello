const R = require('ramda');
import React, { useState, useEffect } from 'react';
import { QueryRenderer } from 'react-relay';
import environment from './index';

const _children = (
  children,
  props
) => {
  if (R.is(Function, children)) return children(props);
  return React.cloneElement(
    children,
    props
  );
};

const Cache = ({
  props,
  renderFetching,
  children
}) => {
  const [$props, setProps] = useState(null);

  useEffect(() => {
    if (!R.isNil(props) && R.isNil($props)) {
      setProps(props);
    } else if (!R.isNil(props) && !R.isNil($props) && !R.equals(props, $props)) {
      setProps(props);
    }
  }, [$props, props]);

  if (R.isNil($props)) return renderFetching;

  return _children(children, $props);
  
};

const Query = ({
  query,
  args,
  renderError,
  renderFetching,
  caching,
  children
}) => {
  return (
    <QueryRenderer
      lookup
      environment={environment()}
      query={query}
      variables={args}
      render={({ error, props }) => {
        if (error) return renderError(error);
        if (!caching) {
          if (R.isNil(props)) return renderFetching;
          return _children(children, props);
        } else {
          return (
            <Cache 
              props={props}
              renderFetching={renderFetching}
            >
              {children}
            </Cache>
          );
        }
      }}
    />
  );
};

Query.defaultProps = {
  renderFetching: <div>fetching</div>,
  renderError: (error) => <div>{error}</div>,
  caching: false
};

export default Query;