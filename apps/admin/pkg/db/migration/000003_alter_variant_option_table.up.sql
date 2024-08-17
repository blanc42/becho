ALTER TABLE variant_options
    RENAME COLUMN name TO value;

ALTER TABLE variant_options
    ADD COLUMN data VARCHAR(255),
    ADD COLUMN image_id CHAR(11);

ALTER TABLE variant_options
    ADD CONSTRAINT fk_variant_option_image
    FOREIGN KEY (image_id) 
    REFERENCES images(id);

-- Update the existing index for the renamed column
DROP INDEX IF EXISTS idx_variant_option_name;
CREATE INDEX idx_variant_option_value ON variant_options(value);