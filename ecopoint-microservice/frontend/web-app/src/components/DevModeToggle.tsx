'use client';

import React from 'react';
import { useDevMode } from '@/contexts/DevModeContext';

const DevModeToggle: React.FC = () => {
  const { isDevMode, toggleDevMode, setMockUser } = useDevMode();

  const handleRoleChange = (role: 'USER' | 'COLLECTOR' | 'ADMIN') => {
    setMockUser({
      uid: `dev-${role.toLowerCase()}-123`,
      email: `dev-${role.toLowerCase()}@ecopoint.com`,
      displayName: `Dev ${role}`,
      role: role,
      photoURL: null,
    });
  };

  return (
    <div className="fixed top-4 right-4 z-50">
      <div className="bg-yellow-100 border border-yellow-400 rounded-lg p-4 shadow-lg">
        <div className="flex items-center space-x-2 mb-2">
          <input
            type="checkbox"
            id="devMode"
            checked={isDevMode}
            onChange={toggleDevMode}
            className="w-4 h-4 text-yellow-600 bg-gray-100 border-gray-300 rounded focus:ring-yellow-500"
          />
          <label htmlFor="devMode" className="text-sm font-medium text-yellow-800">
            🚀 Dev Mode
          </label>
        </div>
        
        {isDevMode && (
          <div className="space-y-2">
            <p className="text-xs text-yellow-700">Switch Role:</p>
            <div className="flex space-x-1">
              <button
                onClick={() => handleRoleChange('USER')}
                className="px-2 py-1 text-xs bg-green-500 text-white rounded hover:bg-green-600"
              >
                User
              </button>
              <button
                onClick={() => handleRoleChange('COLLECTOR')}
                className="px-2 py-1 text-xs bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                Collector
              </button>
              <button
                onClick={() => handleRoleChange('ADMIN')}
                className="px-2 py-1 text-xs bg-purple-500 text-white rounded hover:bg-purple-600"
              >
                Admin
              </button>
            </div>
            <p className="text-xs text-yellow-600">
              Bypass Firebase Auth - Use any email/password to login
            </p>
          </div>
        )}
      </div>
    </div>
  );
};

export default DevModeToggle;