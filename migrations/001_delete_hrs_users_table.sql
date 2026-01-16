-- SQL Script: ลบ HRS_Users table เดิม (ไม่ใช้แล้ว)

USE info_psth
GO

-- ลบ HRS_Users table ถ้าเคยมี
IF EXISTS (SELECT * FROM sys.tables WHERE name = 'HRS_Users' AND type = 'U')
BEGIN
    DROP TABLE HRS_Users;
    PRINT 'Dropped HRS_Users table (no longer needed)';
END
GO

PRINT 'Migration 001: HRS_Users cleanup complete';
GO
