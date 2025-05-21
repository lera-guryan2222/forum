const handleSubmit = async (e) => {
    e.preventDefault();
    try {
        const response = await axios.post(
            'http://localhost:8081/api/v1/posts',
            {
                title: title,    // lowercase
                content: content // lowercase
            },
            {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${authToken}`
                }
            }
        );
        console.log("Full response:", response);
        // ... остальной код ...
    } catch (err) {
        console.error("Full error:", err);
        console.error("Response data:", err.response?.data);
        // ... остальной код ...
    }
};