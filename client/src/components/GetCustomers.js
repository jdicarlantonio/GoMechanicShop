import React, { Fragment, useState } from 'react';

const Customers = () => {
    const [lname, setLname] = useState("");
    const [customers, setCustomers] = useState([]);

    const onSubmitForm = async (e) => {
        e.preventDefault();

        try {
            const body = {
                lname
            };

            const response = await fetch("http://localhost:8080/getCustomersByLastName", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body)
            });
            const jsonData = await response.json();

            if(jsonData.hasOwnProperty('Message')) {
                window.alert(jsonData.Message);
            } else {
                setCustomers(jsonData);
            }
        } catch (error) {
            console.log(error.message);
        }
    }

    return (
        <>
            <Fragment>
                <h1 className="text-center mt-5">Find Customer</h1>
                <form className="mt-4" onSubmit={onSubmitForm}>
                    <input
                        type="text"
                        placeholder="Last Name"
                        className="form-control"
                        value={lname}
                        onChange={
                            e => setLname(e.target.value)
                        }
                    />
                    <button className="btn btn-success">Find</button>
                </form>
            </Fragment>
            <Fragment>
                <table className="table mt-5 text-center">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>First Name</th>
                            <th>Last Name</th>
                            <th>Phone</th>
                            <th>Address</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            customers.map(customer => (
                                <tr key={customer.id}>
                                    <td>{customer.id}</td>
                                    <td>{customer.fname}</td>
                                    <td>{customer.lname}</td>
                                    <td>{customer.phone}</td>
                                    <td>{customer.address}</td>
                                </tr>
                            ))
                        }
                    </tbody>
                </table>
            </Fragment>
        </>
    );
}

export default Customers;