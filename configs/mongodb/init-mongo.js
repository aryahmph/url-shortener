db.createUser({
    user: "aryahmph",
    pwd: "aryahmph",
    roles: [
        {
            role: "readWrite",
            db: "url_shortener",
        },
    ],
});

db = new Mongo().getDB("url_shortener");

db.createCollection("links2", {
    validator: {
        $jsonSchema: {
            bsonType: "object",
            title: "Link Object Validation",
            required: ["original_url"],
            properties: {
                "original_url": {
                    bsonType: "string",
                    description: "'OriginalURL' must be a string and is required"
                }
            }
        }
    }
});
