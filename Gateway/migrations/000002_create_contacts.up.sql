CREATE TABLE contacts (
                          id UUID PRIMARY KEY,
                          user_id UUID NOT NULL,
                          contact_id UUID NOT NULL,
                          status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'accepted', 'blocked')),
                          created_at TIMESTAMP NOT NULL DEFAULT NOW(),

                          CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                          CONSTRAINT fk_contact FOREIGN KEY (contact_id) REFERENCES users(id) ON DELETE CASCADE,
                          CONSTRAINT unique_pair UNIQUE (user_id, contact_id)
);

