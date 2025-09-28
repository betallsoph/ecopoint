'use client';

import React, { createContext, useContext, useEffect, useState } from 'react';
import { 
  User as FirebaseUser,
  onAuthStateChanged,
  signInWithEmailAndPassword,
  createUserWithEmailAndPassword,
  signOut,
  updateProfile
} from 'firebase/auth';
import { auth } from '@/lib/firebase';
import { useDevMode } from './DevModeContext';

interface User {
  uid: string;
  email: string | null;
  displayName: string | null;
  photoURL: string | null;
  role?: 'USER' | 'COLLECTOR' | 'ADMIN';
}

interface AuthContextType {
  user: User | null;
  loading: boolean;
  signIn: (email: string, password: string) => Promise<void>;
  signUp: (email: string, password: string, displayName: string) => Promise<void>;
  logout: () => Promise<void>;
  updateUserProfile: (data: { displayName?: string; photoURL?: string }) => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const { isDevMode, mockUser } = useDevMode();

  useEffect(() => {
    if (isDevMode && mockUser) {
      // Use mock user in dev mode
      const userData: User = {
        uid: mockUser.uid as string,
        email: mockUser.email as string | null,
        displayName: mockUser.displayName as string | null,
        photoURL: mockUser.photoURL as string | null,
        role: mockUser.role as 'USER' | 'COLLECTOR' | 'ADMIN' || 'USER',
      };
      setUser(userData);
      setLoading(false);
      return;
    }

    const unsubscribe = onAuthStateChanged(auth, async (firebaseUser: FirebaseUser | null) => {
      if (firebaseUser) {
        // TODO: Fetch user role from your backend
        const userData: User = {
          uid: firebaseUser.uid,
          email: firebaseUser.email,
          displayName: firebaseUser.displayName,
          photoURL: firebaseUser.photoURL,
          role: 'USER' // Default role, should be fetched from backend
        };
        setUser(userData);
      } else {
        setUser(null);
      }
      setLoading(false);
    });

    return () => unsubscribe();
  }, [isDevMode, mockUser]);

  const signIn = async (email: string, password: string) => {
    setLoading(true);
    try {
      if (isDevMode) {
        // Mock sign in for dev mode
        if (mockUser) {
          const userData: User = {
            uid: mockUser.uid as string,
            email: mockUser.email as string | null,
            displayName: mockUser.displayName as string | null,
            photoURL: mockUser.photoURL as string | null,
            role: mockUser.role as 'USER' | 'COLLECTOR' | 'ADMIN' || 'USER',
          };
          setUser(userData);
        }
        return;
      }
      await signInWithEmailAndPassword(auth, email, password);
    } catch (error) {
      console.error('Sign in error:', error);
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const signUp = async (email: string, password: string, displayName: string) => {
    setLoading(true);
    try {
      const { user: firebaseUser } = await createUserWithEmailAndPassword(auth, email, password);
      await updateProfile(firebaseUser, { displayName });
      
      // TODO: Create user in your backend with the Firebase UID
    } catch (error) {
      console.error('Sign up error:', error);
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const logout = async () => {
    setLoading(true);
    try {
      if (isDevMode) {
        // Mock logout for dev mode
        setUser(null);
        return;
      }
      await signOut(auth);
    } catch (error) {
      console.error('Logout error:', error);
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const updateUserProfile = async (data: { displayName?: string; photoURL?: string }) => {
    if (!auth.currentUser) throw new Error('No user logged in');
    
    try {
      await updateProfile(auth.currentUser, data);
      // TODO: Update user in your backend
    } catch (error) {
      console.error('Update profile error:', error);
      throw error;
    }
  };

  const value = {
    user,
    loading,
    signIn,
    signUp,
    logout,
    updateUserProfile,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}