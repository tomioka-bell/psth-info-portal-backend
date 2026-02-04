-- Migration: Create welfare_benefits table
-- Description: Table to store employee welfare and benefit information with image/file support

IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[welfare_benefits]') AND type in (N'U'))
BEGIN
    CREATE TABLE [dbo].[welfare_benefits] (
        [welfare_benefit_id] INT PRIMARY KEY IDENTITY(1,1),
        [title] NVARCHAR(255) NOT NULL,
        [description] NVARCHAR(MAX) NOT NULL,
        [image_url] NVARCHAR(255),
        [file_name] NVARCHAR(MAX),
        [category] VARCHAR(50) NOT NULL,
        [created_at] DATETIME2 NOT NULL DEFAULT GETUTCDATE(),
        [updated_at] DATETIME2 NOT NULL DEFAULT GETUTCDATE(),
        [deleted_at] DATETIME2 NULL
    );

    -- Create index for soft delete queries
    CREATE NONCLUSTERED INDEX [IX_welfare_benefits_deleted_at] ON [dbo].[welfare_benefits] ([deleted_at]);

    -- Create index for category queries
    CREATE NONCLUSTERED INDEX [IX_welfare_benefits_category] ON [dbo].[welfare_benefits] ([category]);

    -- Create index for search queries
    CREATE NONCLUSTERED INDEX [IX_welfare_benefits_title] ON [dbo].[welfare_benefits] ([title]);

    PRINT 'Table welfare_benefits created successfully'
END
