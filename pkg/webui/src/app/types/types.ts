// Instance describes a Bhojpur Application sidecar instance information
export interface Instance {
    appID: string;
    httpPort: number;
    grpcPort: number;
    appPort: number;
    command: string;
    age: string;
    created: string;
    pid: number;
    replicas: number;
    address: string;
    supportsDeletion: boolean;
    supportsLogs: boolean;
    manifest: string;
    status: string;
    labels: string;
    selector: string;
    config: string;
}

// Status represents the status of a named Bhojpur Application resource
export interface Status {
    service: string;
    name: string;
    namespace: string;
    healthy: string;
    status: string;
    version: string;
    age: string;
    created: string;
}

// Metadata represents metadata from Bhojpur Application sidecar.
export interface Metadata {
    id: string;
    actors: MetadataActors[];
    extended: {[key: string]: any};
}

// MetadataActors represents actor metadata: type and count
export interface MetadataActors {
    type: string;
    count: number;
}

// Log represents a log object supporting log metadata
export interface Log {
    level: string;
    timestamp: number;
    container: string;
    content: string;
}

// AppComponent describes an Bhojpur Application component type
export interface AppComponent {
    name: string;
    kind: string;
    type: string;
    created: string;
    age: string;
    scopes: string[];
    manifest: any;
    img: string;
}

// AppConfiguration represents a Bhojpur Application configuration
export interface AppConfiguration {
    name: string;
    kind: string;
    created: string;
    age: string;
    tracingEnabled: boolean;
    samplingRate: string;
    metricsEnabled: boolean;
    mtlsEnabled: boolean;
    mtlsWorkloadTTL: string;
    mtlsClockSkew: string;
    manifest: any;
}

// YamlViewerOptions describes an options object for the NgMonacoEditor component
export interface YamlViewerOptions {
    folding: boolean;
    minimap: { enabled: boolean };
    readOnly?: boolean;
    language: string;
    theme: string;
    contextmenu?: boolean;
    scrollBeyondLastLine?: boolean;
    lineNumbers?: boolean | any;
}