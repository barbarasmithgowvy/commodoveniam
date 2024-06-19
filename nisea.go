	task := &datastore.Entity{
		Key: datastore.NameKey("Task", "sampletask", nil),
		Properties: []*datastore.Property{
			{
				Name:  "category",
				Value: "Personal",
			},
			{
				Name:  "done",
				Value: false,
			},
			{
				Name:  "priority",
				Value: 4,
			},
			{
				Name:  "description",
				Value: "Learn Cloud Datastore",
			},
		},
	}  
