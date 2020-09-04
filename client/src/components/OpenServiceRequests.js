import React, { Fragment, useEffect, useState } from 'react';

const OpenServiceRequests = () => {
    const [openServiceRequests, setOpenServiceRequests] = useState([]);
    const [message, setMessage] = useState("");

    const getOpenServiceRequests = async () => {
        try {
            const response = await fetch("http://localhost:8080/getOpenServiceRequests");
            const jsonData = await response.json();

            if(jsonData.hasOwnProperty('Message')) {
                setMessage(jsonData.Message);
            } else {
                console.log(jsonData);
                setOpenServiceRequests(jsonData)
            }
        } catch (error) {
            console.log(error.message);
        }
    }

    useEffect(() => {
        getOpenServiceRequests();
    }, []);

    return (
        <Fragment>
            <h1>Open Service Requests</h1>
            <table className="table mt-5 text-center">
                <thead>
                    <tr>
                        <th>Request ID</th>
                        <th>Customer ID</th>
                        <th>Car VIN #</th>
                        <th>Date Opened</th>
                        <th>Odometer Reading</th>
                        <th>Complaint Given</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        openServiceRequests.map(openServiceRequest => (
                            <tr key={openServiceRequest.rid}>
                                <td>{openServiceRequest.rid}</td>
                                <td>{openServiceRequest.customer_id}</td>
                                <td>{openServiceRequest.car_vin}</td>
                                <td>{openServiceRequest.date}</td>
                                <td>{openServiceRequest.odometer}</td>
                                <td>{openServiceRequest.complain}</td>
                            </tr>
                        ))
                    }
                </tbody>
            </table>
            <h5 className="mt-5 text-center">{message}</h5>
        </Fragment>
    );
}

export default OpenServiceRequests;