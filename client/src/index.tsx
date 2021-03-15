import React, {useState} from 'react';
import ReactDOM from 'react-dom';
import {Api} from './swagger/myApi';

const App = () => {
  const [message, setMessage] = useState('')

  const getHello = async () => {
    const api = new Api()
    const response = await api.hello.getHello({})
    const message = response.data.message
    if (message) setMessage(message)
  }

  const postHello = async () => {
    const api = new Api()
    const response = await api.hello.postHello({name: 'Alice'})
    const message = response.data.message
    if (message) setMessage(message)
  }

  return (<>
        <div>{message}</div>
        <button onClick={getHello}>Hello API(GET)</button>
        <button onClick={postHello}>Hello API(POST)</button>
      </>
  )
}

ReactDOM.render(
    <React.StrictMode>
      <App/>
    </React.StrictMode>,
    document.getElementById('root')
);
