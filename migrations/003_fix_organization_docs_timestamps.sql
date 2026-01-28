-- Fix organization_docs table timestamps from bigint to datetime
-- Drop the existing table if it exists with wrong schema
IF OBJECT_ID('organization_docs', 'U') IS NOT NULL
BEGIN
    DROP TABLE organization_docs;
END;

-- Create organization_docs table with correct schema
CREATE TABLE organization_docs (
    organization_doc_id INT IDENTITY(1,1) PRIMARY KEY,
    name NVARCHAR(255) NOT NULL,
    desc NVARCHAR(MAX) NOT NULL,
    department NVARCHAR(50) NOT NULL,
    file_name NVARCHAR(MAX),
    created_at DATETIMEOFFSET DEFAULT GETDATE(),
    updated_at DATETIMEOFFSET DEFAULT GETDATE(),
    deleted_at DATETIMEOFFSET NULL
);

-- Create index on deleted_at for soft deletes
CREATE INDEX IX_organization_docs_deleted_at ON organization_docs(deleted_at);
