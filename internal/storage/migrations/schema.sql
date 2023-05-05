-- Create table for storing data_dumps --
CREATE TABLE IF NOT EXISTS data_dumps (
  id SERIAL NOT NULL PRIMARY KEY,
  sensor TEXT NOT NULL,
  data_values json NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create function for keeping `updated_at` updated with current timestamp
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
  RETURNS trigger AS
$$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$
  LANGUAGE 'plpgsql';

-- Create trigger to call the function we created to keep updated_at updated
DROP trigger IF EXISTS data_dumps_updated_at ON data_dumps;
CREATE trigger data_dumps_updated_at
  BEFORE UPDATE ON data_dumps FOR EACH ROW
  EXECUTE PROCEDURE trigger_set_timestamp();
