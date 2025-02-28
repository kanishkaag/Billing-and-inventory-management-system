import React, { useState, useEffect } from 'react';

function Bills() {
    // State to store the fetched data
    const [bills, setBills] = useState([]);

    // Fetch data when component mounts
    useEffect(() => {
        // Fetching the data from the backend API
        fetch('http://localhost:8080/backend/get_bills') // This is your API URL
            .then(response => response.json())  // Convert the response to JSON
            .then(data => {
                setBills(data); // Set the data in the state
            })
            .catch(error => console.error('Error fetching bills:', error)); // Handle errors
    }, []); // Empty dependency array means this runs once when the component mounts

    return (
        <div>
            <h1>Bills</h1>
            <table border="1">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Customer ID</th>
                    <th>Outlet ID</th>
                    <th>Total Amount</th>
                    <th>Timestamp</th>
                </tr>
                </thead>
                <tbody>
                {bills.map((bill) => (
                    <tr key={bill.id}>
                        <td>{bill.id}</td>
                        <td>{bill.customer_id}</td>
                        <td>{bill.outlet_id}</td>
                        <td>{bill.total_amount}</td>
                        <td>{new Date(bill.time_stamp).toLocaleString()}</td> {/* Format timestamp */}
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
}

export default Bills;
