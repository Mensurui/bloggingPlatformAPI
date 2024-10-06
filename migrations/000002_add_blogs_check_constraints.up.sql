ALTER TABLE blog
ADD CONSTRAINT check_content_length CHECK (length(content) > 10);
