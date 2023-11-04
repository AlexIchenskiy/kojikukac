import { useState, useEffect } from 'react';
import axios from 'axios';

const useFetch = (url, method = 'GET') => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        let response;
        if (method === 'GET') {
          response = await axios.get(url);
        } else if (method === 'POST') {
          response = await axios.post(url);
        } else if (method === 'DELETE') {
          response = await axios.delete(url);
        } else {
          throw new Error('Invalid method. Supported methods are GET, POST, and DELETE.');
        }
        setData(response.data);
        setError(null);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [url, method]);

  return { data, loading, error };
};

export default useFetch;
