import type {Collaborator} from "~/types/user";
import {newRequest} from "~/scripts/request";
import type {Analytics} from "~/types/analytics";

export const fetchAnalytics = async (): Promise<Analytics | null> => {
    try {
        const res = await newRequest(`/admin/analytics`);
        if (res.status != 200) return null;
        const data = await res.json();
        return data as Analytics;
    } catch (e: unknown) {
        console.error(e);
        return null;
    }
};

