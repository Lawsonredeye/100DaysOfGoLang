import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import Form from './welcome'
// import { LiveMsg } from './welcome'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Form />
    {/* <LiveMsg headerMsg = "Hello, Lawson"/> */}
  </StrictMode>,
)
