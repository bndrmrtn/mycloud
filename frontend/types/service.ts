export interface Service {
  service: {
    version: string;
  };
  application: {
    authorization: {
      use_whitelist: boolean;
      use_blacklist: boolean;
      admin: {
        enable_multi_admin: boolean;
      };
    };
  };
}

export interface Pagination<T> {
  next_cursor: string;
  prev_cursor: string;
  data: Array<T>;
}
