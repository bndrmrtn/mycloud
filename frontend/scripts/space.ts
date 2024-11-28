import type { Space, SpaceFile } from '~/types/space';
import { newRequest } from '~/scripts/request';
import { useToast } from 'vue-toastification';

export const space = async (spaceName: string): Promise<Space | Error> => {
  try {
    const res = await newRequest('/spaces', {
      method: 'POST',
      body: JSON.stringify({ name: spaceName }),
    });
    const data = await res.json();

    if (res.status != 200) {
      if (data?.error) return Error(data.error);
      return Error('An unknown error occurred');
    }

    return data as Space;
  } catch (e: unknown) {
    console.error('Failed to create space', e);
    return Error('An unknown error occurred');
  }
};

export const createFile = async (
  id: string,
  file: File,
  dir: string,
  name: string
): Promise<Error | null> => {
  try {
    const data = new FormData();
    data.append('file', file);
    data.append('directory', dir);
    if (name) data.append('filename', name);

    const res = await newRequest(`/spaces/${id}/upload`, {
      method: 'POST',
      body: data,
    });

    if (res.status == 200) return null;

    const err = await res.json();
    if (err?.error) return Error(err.error);
    return Error('An unknown error occurred');
  } catch (e: unknown) {
    console.error('Failed to upload file', e);
    return Error('An unknown error occurred');
  }
};

export const requestSpaceDownload = async (
  spaceID: string,
  path: string = '/'
): Promise<Error | null> => {
  try {
    const res = await newRequest(`/spaces/${spaceID}/download?path=${path}`);
    if (res.status != 202) return Error('Failed to request a download');
    return null;
  } catch (e: unknown) {
    console.error('Failed to fetch files', e);
    return Error('An unknown error occurred');
  }
};

export const fetchSpaces = async (): Promise<Array<Space> | null> => {
  try {
    const res = await newRequest('/spaces');
    if (res.status != 200) return null;
    const data = await res.json();
    return data as Array<Space>;
  } catch (e: unknown) {
    console.error(e);
    return null;
  }
};

export const fetchSpace = async (spaceID: string): Promise<Space | null> => {
  try {
    const res = await newRequest(`/spaces/${spaceID}`);
    if (res.status != 200) return null;
    const data = await res.json();
    return data as Space;
  } catch (e: unknown) {
    console.error(e);
    return null;
  }
};

export const fetchDirs = async (
  spaceID: string,
  path: string
): Promise<Array<string> | null> => {
  try {
    const res = await newRequest(`/spaces/${spaceID}/fs?path=${path}`);
    if (res.status != 200) return null;
    return (await res.json()) as Array<string>;
  } catch (e: unknown) {
    console.error('Failed to fetch dirs', e);
    return null;
  }
};

export const fetchFiles = async (
  spaceID: string,
  path: string
): Promise<Array<SpaceFile> | null> => {
  try {
    const res = await newRequest(`/spaces/${spaceID}/files?path=${path}`);
    if (res.status != 200) return null;
    return (await res.json()) as Array<SpaceFile>;
  } catch (e: unknown) {
    console.error('Failed to fetch files', e);
    return null;
  }
};

export const deleteFile = async (fileID: string): Promise<Error | null> => {
  try {
    const res = await newRequest(`/files/${fileID}`, { method: 'DELETE' });
    if (res.status != 200) return Error('Failed to delete file');
    return null;
  } catch (e: unknown) {
    console.error('Failed to delete file', e);
    return Error('An unknown error occurred');
  }
};

export const updateFileInfo = async (
  fileID: string,
  name: string,
  dir: string
): Promise<Error | null> => {
  try {
    const res = await newRequest(`/files/${fileID}`, {
      method: 'PUT',
      body: JSON.stringify({ name, dir }),
    });
    const data = await res.json();

    if (res.status != 200) {
      if (data?.error) return Error(data.error);
      return Error('An unknown error occurred');
    }

    return null;
  } catch (e: unknown) {
    console.error('Failed to update file', e);
    return Error('An unknown error occurred');
  }
};
