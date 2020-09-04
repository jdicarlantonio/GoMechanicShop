import React, { Fragment, useState } from 'react';

const AddCustomer = () => {
    const [fname, setFname] = useState("");
    const [lname, setLname] = useState("");
    const [phone, setPhone] = useState("");
    const [address, setAddress] = useState("");

    const onSubmitForm = async (e) => {
        e.preventDefault();

        try {
            const body = {
                fname,
                lname,
                phone,
                address
            };

            const response = await fetch("http://localhost:8080/addCustomer", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body)
            });

            console.log(response);
            window.location = "/addCustomer";
        } catch (error) {
            console.error(error.message);
        }
    }

    return (
        <Fragment>
            <h1 className="text-center mt-5">Customer Information</h1>
            <form className="mt-4" onSubmit={onSubmitForm}>
                <input 
                    type="text"
                    placeholder="First Name" 
                    className="form-control" 
                    value={fname} 
                    onChange={
                        e => setFname(e.target.value)
                    }
                />
                <input 
                    type="text" 
                    placeholder="Last Name"  
                    className="form-control"
                    value={lname} 
                    onChange={
                        e => setLname(e.target.value)
                    }
                />
                <input 
                    type="text" 
                    placeholder="Phone" 
                    className="form-control" 
                    value={phone} 
                    onChange={
                        e => setPhone(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Address" 
                    className="form-control" 
                    value={address} 
                    onChange={
                        e => setAddress(e.target.value)
                    }
                />
                <button className="btn btn-success">Add</button>
            </form>
        </Fragment>
    );
}

export default AddCustomer;