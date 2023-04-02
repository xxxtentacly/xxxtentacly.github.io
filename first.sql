CREATE TABLE [users] (
  [email] string,
  [username] nvarchar(255),
  [password] string,
  [id] int,
  [role] nvarchar(255),
  [created_at] timestamp,
  [token] string,
  [bio] text
)
GO

CREATE TABLE [orders] (
  [id] int,
  [user_id] int,
  [status] nvarchar(255),
  [created_at] timestamp
)
GO

CREATE TABLE [order_item] (
  [order_id] int,
  [item_id] int,
  [item_name] text,
  [item_quantity] int
)
GO

ALTER TABLE [users] ADD FOREIGN KEY ([id]) REFERENCES [orders] ([user_id])
GO

ALTER TABLE [orders] ADD FOREIGN KEY ([id]) REFERENCES [order_item] ([order_id])
GO
