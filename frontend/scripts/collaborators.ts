import type {Space} from "~/types/space";
import {newRequest} from "~/scripts/request";
import type {Collaborator} from "~/types/user";

export const fetchCollaborators = async (spaceID: string): Promise<Array<Collaborator> | null> => {
    try {
        const res = await newRequest(`/spaces/${spaceID}/collaborators`);
        if (res.status != 200) return null;
        const data = await res.json();
        return data as Array<Collaborator>;
    } catch (e: unknown) {
        console.error(e);
        return null;
    }
};

export const putCollaborator = async (spaceID: string, email: string, permission: {create: boolean, read: boolean, update: boolean, delete: boolean}): Promise<Error|null> => {
    try {
        const res = await newRequest(`/spaces/${spaceID}/collaborators`, {
            method: 'PUT',
            body: JSON.stringify({email, permission}),
        });
        if (res.status == 200) return null

        const err = await res.json();
        if (err?.error) return Error(err.error);
        return Error('An unknown error occurred');
    } catch (e: unknown) {
        console.error(e);
        return Error('Failed to modify collaborator');
    }
}