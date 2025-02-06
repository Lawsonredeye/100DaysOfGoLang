import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {Person} from './welcome'
// import { LiveMsg } from './welcome'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Person />
    {/* <LiveMsg headerMsg = "Hello, Lawson"/> */}
  </StrictMode>,
)
