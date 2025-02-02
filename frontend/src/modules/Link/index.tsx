import React, { useEffect, useState } from 'react';
import { usePlaidLink } from 'react-plaid-link';
import { useGlobalState } from '../../GlobalState';


const PlaidIntegration: React.FC = () => {
  const [linkToken, setLinkToken] = useState<string | null>(null);

  useEffect(() => {
    const generateToken = async () => {
      const response = await fetch('/api/createLinkToken', {
        method: 'POST',
      });
      const data = await response.json();
      setLinkToken(data.link_token);
    };

    generateToken();
    console.log(linkToken)
  }, []);

  return linkToken ? <PlaidLinkHandler linkToken={linkToken} /> : <p>Loading...</p>;
};

const PlaidLinkHandler: React.FC<{ linkToken: string }> = ({ linkToken }) => {
  const {credentials} = useGlobalState()
  const onSuccess = (public_token: string) => {
    if (credentials != undefined) {
    fetch('/api/getAccessToken', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        credentials,
        public_token }),
    });
    }
  };

  const { open, ready } = usePlaidLink({ token: linkToken, onSuccess });

  useEffect(() => {
    if (ready) {
      open(); // Open Plaid Link automatically when page loads
    }
  }, [ready, open]);

  return null; // No UI, just auto-triggers the flow
};

export default PlaidIntegration;
