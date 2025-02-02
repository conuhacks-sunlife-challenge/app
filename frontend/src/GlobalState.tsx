import { createContext, useContext, useState, ReactNode } from "react";
import Credentials from "./types/credentials";

type GlobalState = {
  credentials: Credentials | undefined,
  setCredentials: (value: Credentials) => void;
  loggedIn: boolean,
  setLoggedIn: (value: boolean) => void;
};

const GlobalStateContext = createContext<GlobalState | undefined>(undefined);

export const useGlobalState = () => {
  const context = useContext(GlobalStateContext);
  if (!context) {
    throw new Error("useGlobalState must be used within a GlobalStateProvider");
  }
  return context;
};

export const GlobalStateProvider = ({ children }: { children: ReactNode }) => {
  const [credentials, setCredentials] = useState<Credentials | undefined>(undefined);
  const [loggedIn, setLoggedIn] = useState<boolean>(false);


  return (
    <GlobalStateContext.Provider value={{ 
      credentials,
      setCredentials,
      loggedIn,
      setLoggedIn,
    }}>
      {children}
    </GlobalStateContext.Provider>
  );
};
