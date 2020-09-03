import React, { Fragment, useState } from 'react';

const CloseServiceRequest = () => {
    const [rid, setRid] = useState("");
    const [mid, setMid] = useState("");
    const [comment, setComment] = useState("");
    const [bill, setBill] = useState("");

    const onSubmitForm = async (e) => {
        e.preventDefault();

        try {
            const body = {
                rid,
                mid,
                comment,
                bill
            };

            const response = await fetch("http://localhost:8080/closeServiceRequest", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(body)
            });

            console.log(response);
            window.location = "/";
        } catch (error) {
            console.log(error.message)
        }
    }

    return (
        <Fragment>
            <h1 className="text-center mt-5">Close Service Request</h1>
            <form className="mt-4" onSubmit={onSubmitForm}>
                <input
                    type="text"
                    placeholder="Request ID"
                    className="form-control"
                    value={rid}
                    onChange={
                        e => setRid(e.target.value)
                    }
                />
                <input
                    type="text"
                    placeholder="Employee ID"
                    className="form-control"
                    value={mid}
                    onChange={
                        e => setMid(e.target.value)
                    }
                />
                <input
                    type="text"
                    placeholder="Employee's Comment"
                    className="form-control"
                    value={comment}
                    onChange={
                        e => setComment(e.target.value)
                    }
                />
                <input
                    type="text"
                    placeholder="Bill Amount (USD)"
                    className="form-control"
                    value={bill}
                    onChange={
                        e => setBill(e.target.value)
                    }
                />
                <button className="btn btn-success">Close</button>
            </form>
        </Fragment>
    );
}

export default CloseServiceRequest;