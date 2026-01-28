-- Migration: Create ps_organizations table
-- Description: Table to store organization departments with category and icon information

IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[ps_organizations]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[ps_organizations] (
        [id] INT PRIMARY KEY IDENTITY(1,1),
        [name] NVARCHAR(255) NOT NULL,
        [desc] NVARCHAR(MAX) NOT NULL,
        [category] VARCHAR(50) NOT NULL,
        [href] NVARCHAR(MAX) NOT NULL,
        [icon] NVARCHAR(MAX),
        [created_at] DATETIME2 NOT NULL DEFAULT GETUTCDATE(),
        [updated_at] DATETIME2 NOT NULL DEFAULT GETUTCDATE(),
        [deleted_at] DATETIME2 NULL
    );

    -- Create index for soft delete queries
    CREATE NONCLUSTERED INDEX [IX_ps_organizations_deleted_at] ON [dbo].[ps_organizations] ([deleted_at]);

    -- Create index for category filtering
    CREATE NONCLUSTERED INDEX [IX_ps_organizations_category] ON [dbo].[ps_organizations] ([category]);

    PRINT 'Table ps_organizations created successfully';
END
ELSE
BEGIN
    PRINT 'Table ps_organizations already exists';
END
