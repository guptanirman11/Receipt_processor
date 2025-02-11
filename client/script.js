
var receiptID;
async function uploadReceipt(){
    const jsonInput = document.getElementById("jsonInput").value.trim()

    if (jsonInput === "") {
        alert("Please paste JSON data.");
        return;
    }

    try {
        // JSON.parse(jsonInput)

        const response = await fetch("http://localhost:8080/receipts/process", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: jsonInput
        });

        const data = await response.json();

        if (!response.ok) {
            document.getElementById("points").innerText = `Error: ${data.error}`;
            return;
        }

        receiptID = data.id
        document.getElementById("receiptId").innerText = receiptID;

        // Automatically fill in the receipt ID input field
        document.getElementById("receiptIDInput").value = receiptID;
        // getPoints(receiptID)

    } catch (error) {

        alert("Invalid JSON");
        document.getElementById("points").innerText = "Invalid JSON format.";
    }


}

async function getPointsFromInput() {
    const receiptID = document.getElementById("receiptIDInput").value.trim();

    if (receiptID === "") {
        alert("Please enter a Receipt ID.");
        return;
    }

    getPoints(receiptID);
}

async function getPoints(receiptID) {
    try {

        const response = await fetch(`http://localhost:8080/receipts/${receiptID}/points`);
        const data = await response.json()

        if (response.ok) {
            document.getElementById("points").innerText = data.points;
        } else {
            document.getElementById("points").innerText = `Error: ${data.error}`;
        }

    } catch (error) {

        console.error("Request failed:", error);
        document.getElementById("points").innerText = "Error fetching points.";

    }
    
}

async function clearInput(){
    document.getElementById("jsonInput").value = "";
}