import { createClient } from "@supabase/supabase-js";
import { queryOptions } from "@tanstack/react-query";
import { PluginStatus } from "./supabase-types";

const supabaseQueryKey = "supabase";
const supabaseUrl = "http://localhost:8000";
const supabaseKey =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyAgCiAgICAicm9sZSI6ICJhbm9uIiwKICAgICJpc3MiOiAic3VwYWJhc2UtZGVtbyIsCiAgICAiaWF0IjogMTY0MTc2OTIwMCwKICAgICJleHAiOiAxNzk5NTM1NjAwCn0.dc_X5iR_VP_qT0zsiyj_I_OZ2T9FtRU2BBNWN8Bu4GE";

export const supabase = createClient(supabaseUrl, supabaseKey);

/**
 * Get persistence plugin status from the database using it's name.
 */
export function persistencePluginStatus(name: string) {
  const tableName = "plugin_status";

  return queryOptions({
    queryKey: [supabaseQueryKey, tableName, name],
    queryFn: async () => {
      const { data, error } = await supabase
        .from(tableName)
        .select()
        .match({ name });

      if (error) {
        throw new Error(
          `Could not get persistence plugin status with name ${name}: ${error.message}`,
        );
      }

      if (data.length !== 1) {
        throw new Error(
          `Invalid response when persistence plugin status with name ${name}: ${data}`,
        );
      }

      return data.pop() as PluginStatus;
    },
  });
}
