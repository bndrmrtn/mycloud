export interface User {
  id: string;
  gid: string;
  name: string;
  email: string;
  image_url: string;
  role: 'user' | 'admin';
  created_at: string;
}

export interface Collaborator {
  id: string
  user: User
  permission: {
    read_file: boolean
    update_file: boolean
    delete_file: boolean
    upload_file: boolean
  }
}