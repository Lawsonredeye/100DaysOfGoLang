import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { Welcome, WelcomeMsg } from './welcome'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    {/* <App /> */}
    <Welcome />
    <WelcomeMsg />
  </StrictMode>,
)
