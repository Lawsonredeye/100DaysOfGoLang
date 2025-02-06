import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {App} from './welcome'
// import { LiveMsg } from './welcome'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
    {/* <LiveMsg headerMsg = "Hello, Lawson"/> */}
  </StrictMode>,
)
