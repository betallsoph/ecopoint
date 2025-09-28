'use client';

import React, { createContext, useContext, useState, ReactNode } from 'react';

interface DevModeContextType {
  isDevMode: boolean;
  toggleDevMode: () => void;
  mockUser: Record<string, unknown> | null;
  setMockUser: (user: Record<string, unknown>) => void;
}

const DevModeContext = createContext<DevModeContextType | undefined>(undefined);

export const useDevMode = () => {
  const context = useContext(DevModeContext);
  if (!context) {
    throw new Error('useDevMode must be used within a DevModeProvider');
  }
  return context;
};

interface DevModeProviderProps {
  children: ReactNode;
}

export const DevModeProvider: React.FC<DevModeProviderProps> = ({ children }) => {
  const [isDevMode, setIsDevMode] = useState(false);
  const [mockUser, setMockUser] = useState<Record<string, unknown> | null>(null);

  const toggleDevMode = () => {
    setIsDevMode(!isDevMode);
    if (!isDevMode) {
      // Set default mock user when enabling dev mode
      setMockUser({
        uid: 'dev-user-123',
        email: 'dev@ecopoint.com',
        displayName: 'Dev User',
        role: 'USER',
        photoURL: null,
      });
    } else {
      setMockUser(null);
    }
  };

  return (
    <DevModeContext.Provider value={{ isDevMode, toggleDevMode, mockUser, setMockUser }}>
      {children}
    </DevModeContext.Provider>
  );
};