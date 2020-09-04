import React, { Fragment, useState } from 'react';

const CreateServiceRequest = () => {
    const [fname, setFname] = useState("");
    const [lname, setLname] = useState("");
    const [phone, setPhone] = useState("");
    const [vin, setVin] = useState("");
    const [odometer, setOdometer] = useState("");
    const [complaint, setComplaint] = useState("");

    const onSubmitForm = async (e) => {
        e.preventDefault();

        try {
            const body = {
                fname,
                lname,
                phone,
                vin, 
                odometer,
                complaint
            };

            const response = await fetch("http://localhost:8080/addServiceRequest", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body)
            });
            const jsonData = await response.json();

            if(jsonData.hasOwnProperty('Message')) {
                window.alert(jsonData.Message);
            } else {
                console.log(response);
                window.location = "/";
            }
        } catch (error) {
            console.log(error.message);
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
                    placeholder="VIN #" 
                    className="form-control" 
                    value={vin} 
                    onChange={
                        e => setVin(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Odometer Reading" 
                    className="form-control" 
                    value={odometer} 
                    onChange={
                        e => setOdometer(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Complaint Given" 
                    className="form-control" 
                    value={complaint} 
                    onChange={
                        e => setComplaint(e.target.value)
                    }
                />
                <button className="btn btn-success">Add</button>
            </form>
        </Fragment>
    );
}

export default CreateServiceRequest;