export interface User {
    login: string;
    id: number;
    avatar_url: string;
    name: string;
    email?: string;
    public_repos: number;
    public_gists: number;
    followers: number;
    following: number;
    created_at: string;
    updated_at: string;
    private_gists: number;
    total_private_repos: number;
    owned_private_repos: number;
    two_factor_authentication: boolean;
    organization: string;
  }
  
export interface Organization {
  id: number;
  login: string;
  avatar_url: string;
  description?: string;
  url: string;
}

export interface Repository {
  name: string;
  description: string;
  updated_at: string;
  private: boolean;
  stargazers_count: number;
  html_url: string;
}

export interface Member {
  id: number;
  login: string;
  avatar_url: string;
  html_url: string;
}
