ALTER TABLE transaction
  ADD COLUMN
    extension_options     JSON    NOT NULL DEFAULT '[]'::JSON;