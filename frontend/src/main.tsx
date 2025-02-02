import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Home from './modules/home/index.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Home />
  </StrictMode>,
)
