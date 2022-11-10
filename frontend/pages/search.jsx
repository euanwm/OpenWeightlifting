import { useState } from 'react';

const SearchLifters = async (
    lifterName = ''
) => {
    const bodyContent = JSON.stringify({lifterName})
    const urlQuery = `${process.env.API}/search?name=` + lifterName

    const res = await fetch(urlQuery, {
        method: "GET",
        headers: {
            Accept: '*/*',
            'Content-Type': 'application/json',
        }
    }).catch(error => console.error(error))

    return await res.json()
}

const SearchFilter = ({  }) => {
    let oldSearch = ""
    const [searchInput, SetSearchInput] = useState("");
    let FilteredData = [{"name": "None"}]
    if (searchInput.length >= 3 && searchInput !== oldSearch) {
        let FilteredData = SearchLifters(searchInput)
        oldSearch = searchInput
    }

    return (
        <>
            <div className="container-fluid mt-4 mb-4">
                <div className="row justify-content-center">
                    <div className="col-md-10">
                        <div className="card">
                            <div className="card-body p-3">
                                <div className="row justify-content-between align-items-center">
                                    <div className="col-md-3">
                                            <h5 className="text-primary">
                                                {FilteredData.length} results found.
                                            </h5>
                                    </div>
                                    <div className="col-md-3">
                                        <div className="form-group mb-0">
                                            <input type="text" className="form-control" placeholder="Search" value={searchInput} onChange={(e) => SetSearchInput(e.target.value)} />
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="card-body p-0">
                                <div className="table-responsive">
                                    <table className="table table-text-small mb-0">
                                        <thead className="thead-dark table-sorting">
                                        <tr>
                                            <th>#</th>
                                            <th>Name</th>
                                        </tr>
                                        </thead>
                                        <tbody>
                                        {FilteredData.map((data, index) => {
                                            const { id, name } = data;
                                            return (
                                                <tr key={index}>
                                                    <td>{id}</td>
                                                    <td>{name}</td>
                                                </tr>
                                            );
                                        })}
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default SearchFilter
