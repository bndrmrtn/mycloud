export interface User {
  id: string;
  gid: string;
  name: string;
  email: string;
  image_url: string;
  role: 'user' | 'admin';
  created_at: string;
}
