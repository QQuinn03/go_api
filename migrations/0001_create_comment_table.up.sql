CREATE TABLE IF NOT EXISTS comment(
    ID uuid,
    Slug text,
    Author text,
    Body text

);

/*psql -u Connect to the database as the user username instead of the default. (You must have permission to do so, of course.)*/