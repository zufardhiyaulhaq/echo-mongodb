db = db.getSiblingDB('echo-mongodb');
db.createUser(
    {
        user: "echo-mongodb",
        pwd: "echo-mongodb",
        roles: [
            {
                role: "readWrite",
                db: "echo-mongodb"
            }
        ]
    }
);
