import { gql } from 'graphql-tag';

export const userTypeDefs = gql`
  type User {
    id: ID!
    firebaseUid: String!
    email: String!
    firstName: String!
    lastName: String!
    phone: String
    role: UserRole!
    isActive: Boolean!
    profileImageUrl: String
    addresses: [Address!]!
    preferences: UserPreferences
    createdAt: String!
    updatedAt: String!
  }

  type Address {
    street: String!
    city: String!
    state: String!
    zipCode: String!
    coordinates: Coordinates!
  }

  type Coordinates {
    lat: Float!
    lng: Float!
  }

  type UserPreferences {
    language: String!
    notifications: Boolean!
    theme: Theme!
  }

  enum UserRole {
    USER
    COLLECTOR
    ADMIN
  }

  enum Theme {
    LIGHT
    DARK
  }

  input CreateUserInput {
    firebaseUid: String!
    email: String!
    firstName: String!
    lastName: String!
    phone: String
    role: UserRole = USER
    profileImageUrl: String
  }

  input UpdateUserInput {
    firstName: String
    lastName: String
    phone: String
    profileImageUrl: String
  }

  input AddressInput {
    street: String!
    city: String!
    state: String!
    zipCode: String!
    coordinates: CoordinatesInput!
  }

  input CoordinatesInput {
    lat: Float!
    lng: Float!
  }

  input UserPreferencesInput {
    language: String
    notifications: Boolean
    theme: Theme
  }

  input FirebaseAuthInput {
    firebaseToken: String!
  }

  type AuthPayload {
    user: User!
  }

  type Query {
    me: User
    users(role: UserRole, isActive: Boolean): [User!]!
    user(id: ID!): User
  }

  type Mutation {
    createUser(input: CreateUserInput!): AuthPayload!
    authenticateUser(input: FirebaseAuthInput!): AuthPayload!
    updateProfile(input: UpdateUserInput!): User!
    deactivateUser(id: ID!): Boolean!
    activateUser(id: ID!): Boolean!
    addAddress(input: AddressInput!): Address!
    updateAddress(id: ID!, input: AddressInput!): Address!
    deleteAddress(id: ID!): Boolean!
    updatePreferences(input: UserPreferencesInput!): UserPreferences!
  }
`;