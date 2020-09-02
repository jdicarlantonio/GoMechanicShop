import React, { Fragment, useState } from 'react';

const AddCar = () => {
    const [vin, setVin] = useState("");
    const [make, setMake] = useState("");
    const [model, setModel] = useState("");
    const [year, setYear] = useState("");

    // car owner information
    const [fname, setFname] = useState("");
    const [lname, setLname] = useState("");
    const [phone, setPhone] = useState("");

    const onSubmitForm = async (e) => {
        e.preventDefault();

        try{
            const body = {
                vin,
                make,
                model,
                year,
                fname,
                lname,
                phone
            };

            const response = await fetch("http://localhost:8080/addCar", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body)
            });

            if(response.status === 204) {
                window.alert("Customer does not exist in database.\nPlease add customer to database.");
            }

            window.location = "/";
        } catch(error) {
            console.error(error.message);
        }
    }

    return (
        <Fragment>
            <h1 className="text-center mt-5">Vehicle Information</h1>
            <form className="mt-4" onSubmit={onSubmitForm}>
                <input 
                    type="text"
                    placeholder="VIN Number" 
                    className="form-control" 
                    value={vin} 
                    onChange={
                        e => setVin(e.target.value)
                    }
                />
                <input 
                    type="text" 
                    placeholder="Make"  
                    className="form-control"
                    value={make} 
                    onChange={
                        e => setMake(e.target.value)
                    }
                />
                <input 
                    type="text" 
                    placeholder="Model" 
                    className="form-control" 
                    value={model} 
                    onChange={
                        e => setModel(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Year" 
                    className="form-control" 
                    value={year} 
                    onChange={
                        e => setYear(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Owner's First Name" 
                    className="form-control" 
                    value={fname} 
                    onChange={
                        e => setFname(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Owner's Last Name" 
                    className="form-control" 
                    value={lname} 
                    onChange={
                        e => setLname(e.target.value)
                    }
                />
                <input 
                    type="text"
                    placeholder="Owner's Phone Number" 
                    className="form-control" 
                    value={phone} 
                    onChange={
                        e => setPhone(e.target.value)
                    }
                />
                <button className="btn btn-success">Add</button>
            </form>
        </Fragment>
    );
}

export default AddCar;