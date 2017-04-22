import { EventEmitter } from "events";
import { hashHistory } from "react-router";
import "whatwg-fetch";
import dispatcher from "../dispatcher";
import tokenStore from "./TokenStore";

var checkGetActivationStatus = (response) => {
  if (response.status >= 200 && response.status < 300) {
    return response
  } else {
    // dont return an error when there is no activation
  }
};

var checkStatus = (response) => {
  if (response.status >= 200 && response.status < 300) {
    return response
  } else {
    throw response.json();
  }
};

var errorHandler = (error) => {
  console.log(error);
  error.then((data) => {
    dispatcher.dispatch({
      type: "CREATE_ERROR",
      error: data,
    });

    if (data.Code === 16) {
      hashHistory.push("/jwt");
    }
  });
};

class NodeStore extends EventEmitter {
  getAll(applicationID, pageSize, offset, callbackFunc) {
    fetch("/api/applications/"+applicationID+"/nodes?limit="+pageSize+"&offset="+offset, {headers: tokenStore.getHeader()})
      .then(checkStatus)
      .then((response) => response.json())
      .then((responseData) => {
        if(typeof(responseData.result) === "undefined") {
          callbackFunc(0, []);
        } else {
          callbackFunc(responseData.totalCount, responseData.result);
        }
      })
      .catch(errorHandler);
  }

  getNode(applicationID, name, callbackFunc) {
    fetch("/api/nodes/"+name, {headers: tokenStore.getHeader()})
      .then(checkStatus)
      .then((response) => response.json())
      .then((responseData) => {
        callbackFunc(responseData);
      })
      .catch(errorHandler);
  }

  createNode(applicationID, node, callbackFunc) {
    fetch("/api/nodes", {method: "POST", body: JSON.stringify(node), headers: tokenStore.getHeader()})
      .then(checkStatus)
      .then((response) => response.json())
      .then((responseData) => {
        callbackFunc(responseData);
      })
      .catch(errorHandler);
  }

  updateNode(applicationID, devEUI, node, callbackFunc) {
    fetch("/api/nodes/"+devEUI, {method: "PUT", body: JSON.stringify(node), headers: tokenStore.getHeader()})
      .then(checkStatus)
      .then((response) => response.json())
      .then((responseData) => {
        callbackFunc(responseData);
      })
      .catch(errorHandler);
  }

  deleteNode(applicationID, devEUI, callbackFunc) {
    fetch("/api/nodes/"+devEUI, {method: "DELETE", headers: tokenStore.getHeader()})
      .then(checkStatus)
      .then((response) => response.json())
      .then((responseData) => {
        callbackFunc(responseData);
      })
      .catch(errorHandler);
  }

  activateNode(applicationID, devEUI, activation, callbackFunc) {
    fetch("/api/nodes/"+devEUI+"/activation", {method: "POST", body: JSON.stringify(activation), headers: tokenStore.getHeader()})
      .then(checkStatus)
      .then((response) => response.json())
      .then((responseData) => {
        callbackFunc(responseData);
      })
      .catch(errorHandler);
  }

  getActivation(applicationID, devEUI, callbackFunc) {
    fetch("/api/nodes/"+devEUI+"/activation", {headers: tokenStore.getHeader()})
      .then(checkGetActivationStatus)
      .then((response) => response.json())
      .then((responseData) => {
        callbackFunc(responseData);
      })
      .catch(errorHandler);
  }
}

const nodeStore = new NodeStore();

export default nodeStore;
