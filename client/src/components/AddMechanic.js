import React, { Fragment, useState } from 'react';

const AddMechanic = () => {
    const [fname, setFname] = useState("");
    const [lname, setLname] = useState("");
    const [experience, setExperience] = useState("");

    const onSubmitForm = async (e) => {
        e.preventDefault();

        try {
            const body = {
                fname,
                lname,
                experience
            };

            const response = await fetch("http://localhost:8080/addMechanic", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body)
            });
            
            console.log(response);
            window.location = "/"
        } catch (error) {
            console.error(error.message);
        }
    }

    return (
        <Fragment>
            <h1 className="text-center mt-5">Mechanic</h1>
            <form className="d-flex mt-5" onSubmit={onSubmitForm}>
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
                    placeholder="Years of Experience"
                    className="form-control"
                    value={experience}
                    onChange={
                        e => setExperience(e.target.value)
                    }
                />
                <button className="btn btn-success">Add</button>
            </form>
        </Fragment>
    );
}

export default AddMechanic;