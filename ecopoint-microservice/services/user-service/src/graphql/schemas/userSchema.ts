import { gql } from 'graphql-tag';

export const userTypeDefs = gql`
  type User {
    id: ID!
    email: String!
    firstName: String!
    lastName: String!
    phone: String!
    role: UserRole!
    isActive: Boolean!
    profileImage: String
    address: Address
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
    email: String!
    password: String!
    firstName: String!
    lastName: String!
    phone: String!
    role: UserRole = USER
    address: AddressInput
    preferences: UserPreferencesInput
  }

  input UpdateUserInput {
    firstName: String
    lastName: String
    phone: String
    address: AddressInput
    preferences: UserPreferencesInput
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

  input LoginInput {
    email: String!
    password: String!
  }

  type AuthPayload {
    token: String!
    user: User!
  }

  type Query {
    me: User
    users(role: UserRole, isActive: Boolean): [User!]!
    user(id: ID!): User
  }

  type Mutation {
    register(input: CreateUserInput!): AuthPayload!
    login(input: LoginInput!): AuthPayload!
    updateProfile(input: UpdateUserInput!): User!
    changePassword(currentPassword: String!, newPassword: String!): Boolean!
    deactivateUser(id: ID!): Boolean!
    activateUser(id: ID!): Boolean!
  }
`;