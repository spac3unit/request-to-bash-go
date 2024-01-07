import { useState } from 'react';
import styles from './App.module.css';

function App() {
  const [cpu, setCpu] = useState('');
  const [ram, setRam] = useState('');
  const [response, setResponse] = useState('');

  const handleSubmit = async (e: any) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8081/create/server', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `cpu=${cpu}&ram=${ram}`,
      });

      const data = await response.json();
      setResponse(data);
    } catch (error) {
      console.error('Error sending POST request:', error);
      setResponse('Error sending POST request');
    }
  };

  return (
    <div className={styles.App}>
      <h1>/create/server Request Example</h1>
      <form className={styles.form} onSubmit={handleSubmit}>
        <label>
          CPU:
          <input type="text" value={cpu} onChange={(e) => setCpu(e.target.value)} />
        </label>
        <br />
        <label>
          RAM:
          <input type="text" value={ram} onChange={(e) => setRam(e.target.value)} />
        </label>
        <br />
        <button type="submit">Submit</button>
      </form>
      <div className={styles.response}>
        <h2>Response:</h2>
        <pre>{JSON.stringify(response, null, 2)}</pre>
      </div>

      <div>
        <p>Curl example:</p>
        <code>curl -X POST -d "cpu=4&ram=8" http://localhost:8081/create/server</code>
      </div>
    </div>
  );
}

export default App;