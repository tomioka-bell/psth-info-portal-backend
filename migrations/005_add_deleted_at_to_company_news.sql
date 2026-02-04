-- Migration: Add deleted_at column to ps_company_news table
-- Description: Add soft delete support to CompanyNews model

IF NOT EXISTS (SELECT * FROM sys.columns WHERE object_id = OBJECT_ID(N'[dbo].[ps_company_news]') AND name = 'deleted_at')
BEGIN
    ALTER TABLE [dbo].[ps_company_news] 
    ADD [deleted_at] DATETIME2 NULL;

    -- Create index for soft delete queries
    CREATE NONCLUSTERED INDEX [IX_ps_company_news_deleted_at] ON [dbo].[ps_company_news] ([deleted_at]);

    PRINT 'Column deleted_at added to ps_company_news table successfully'
END
