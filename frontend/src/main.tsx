import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import AppWrapper from './App.tsx'
import { GlobalStateProvider } from './GlobalState.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <GlobalStateProvider>
      <AppWrapper />
    </GlobalStateProvider>
  </StrictMode>,
)
