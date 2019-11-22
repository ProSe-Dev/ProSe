import React from 'react';
import './ProjectPage.css'
import TitleBar from 'components/TitleBar';
import Table from 'components/Table';
import ToggleSwitch from 'components/ToggleSwitch';

function SnapshotOutdatedAlert() {
  return (
    <div class="alert alert-warning projectpage-outdated-alert">
      <div class="vcenter-wrapper">
        Hey! Your project snapshot is outdated
      </div>
      <div class="hleft-wrapper">
        <button type="button" class="btn btn-warning">Renew</button>
      </div>
    </div>
  );
}

function SnapshotUptodateAlert() {
  return (
    <div class="alert alert-primary">
      Everything is up to date!
    </div>
  );
}

function Title(props){
  function Dot(status){
    if (status.synced)
      return (<span class="green dot"></span>);
    else
      return (<span class="red dot"></span>);
  }
  return(
    <span class="title-container">
      <span class="display-3 font-weight-bold project-title"> {props.ProjectName} </span>
      <Dot synced = {props.CertStatus}/>
      <button class="btn btn-primary btn-lg cert-button">New Certificate</button>
    </span>
  )
}

const UpToDateStatus = (<span className="badge badge-pill badge-success">uptodate</span>);
const OutdatedStatus = (<span className="badge badge-pill badge-danger">outdated</span>);
const ExcludedStatus = (<span className="badge badge-pill badge-secondary">excluded</span>);
const FILE_TABLE_HEADERS = ['Status', 'File Name', 'Include In Snapshots', 'Other Actions'];
const test_rows = [
  [UpToDateStatus, 'Cover Page.pdf', <ToggleSwitch />, 'xxx'],
  [OutdatedStatus, 'Cover Page.pdf', <ToggleSwitch />, 'xxx'],
  [ExcludedStatus, 'Cover Page.pdf', <ToggleSwitch />, 'xxx'],
  [ExcludedStatus, 'Cover Page.pdf', <ToggleSwitch />, 'xxx'],
  [ExcludedStatus, 'Cover Page.pdf', <ToggleSwitch />, 'xxx']
];

class Files extends React.Component{
  render (){
    return(
      <div>
        <div className="projectpage-heading">
          Project Files
        </div>
        <Table
          headers={FILE_TABLE_HEADERS}
          rows={test_rows}
        />
      </div>
    );
  }
}

class Certificates extends React.Component {
  
  render(){
    const cert_list = getCerts();
    return(
      <div class="certificates shadow">
        <div class="display-4">
          Certificates
        </div>
        <div class="cert-table">
          {cert_list}
        </div>
    </div>
      
    )
  }
}

class Summary extends React.Component {
  render(){
    return(
      <div class ="summary shadow">
        <div class="display-4">
            Summary
        </div>
        <p>{this.props.description}</p>
    </div>
    )
  }
}

class ProjectPage extends React.Component {
  render() {
    return (
      <div class ="main-container">
        <div class ="inner-container">
          <TitleBar 
            title="hello"
            subtitle="bye"
            showSettings
          />
          <SnapshotOutdatedAlert />
          <Files
            FileList = {"temp"}
          />
          <Certificates
            CertList = {"temp"}
          />
        </div>
      </div>
    );
  }
}

export default ProjectPage;
/** creates a html table from list of files */
function getFiles(){
  return(
    <table class="table">
        <thead>
          <tr class = "file-row-dark">
            <th scope="col">#</th>
            <th scope="col">First</th>
            <th scope="col">Last</th>
            <th scope="col">Handle</th>
          </tr>
        </thead>
        <tbody>
          <tr class="file-row-light">
            <th scope="row">1</th>
            <td>Mark</td>
            <td>Otto</td>
            <td>@mdo</td>
          </tr>
          <tr class="file-row-light">
            <th scope="row">2</th>
            <td>Jacob</td>
            <td>Thornton</td>
            <td>@fat</td>
          </tr>
          <tr class="file-row-light">
            <th scope="row">3</th>
            <td>Larry</td>
            <td>the Bird</td>
            <td>@twitter</td>
          </tr>
        </tbody>
      </table>
  )
}
/** creates a html table from list of certificates */
function getCerts(){
  return (<table class="table">
  <thead>
    <tr class="bg-warning">
      <th scope="col">#</th>
      <th scope="col">First</th>
      <th scope="col">Last</th>
      <th scope="col">Handle</th>
    </tr>
  </thead>
  <tbody>
    <tr class="table-light">
      <th scope="row">1</th>
      <td>Mark</td>
      <td>Otto</td>
      <td>@mdo</td>
    </tr>
    <tr class="table-light">
      <th scope="row">2</th>
      <td>Jacob</td>
      <td>Thornton</td>
      <td>@fat</td>
    </tr>
    <tr class="table-light">
      <th scope="row">3</th>
      <td>Larry</td>
      <td>the Bird</td>
      <td>@twitter</td>
    </tr>
    </tbody>
  </table>)
}