/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  "/health": {
    get: {
      responses: {
        /** The service is healthy */
        204: never;
      };
    };
  };
  "/sync": {
    post: {
      responses: {
        /** The time and metadata is synced with the host */
        204: never;
      };
    };
  };
  "/files": {
    get: {
      parameters: {
        query: {
          /** Path to the file, URL encoded. Can be relative to user's home directory. */
          path?: components["parameters"]["FilePath"];
          /** User used for setting the owner, or resolving relative paths. */
          username: components["parameters"]["User"];
        };
      };
      responses: {
        200: components["responses"]["DownloadSuccess"];
        400: components["responses"]["DirectoryPathError"];
        401: components["responses"]["InvalidUser"];
        404: components["responses"]["FileNotFound"];
        500: components["responses"]["InternalServerError"];
      };
    };
    post: {
      parameters: {
        query: {
          /** Path to the file, URL encoded. Can be relative to user's home directory. */
          path?: components["parameters"]["FilePath"];
          /** User used for setting the owner, or resolving relative paths. */
          username: components["parameters"]["User"];
        };
      };
      responses: {
        204: components["responses"]["UploadSuccess"];
        400: components["responses"]["DirectoryPathError"];
        401: components["responses"]["InvalidUser"];
        500: components["responses"]["InternalServerError"];
        507: components["responses"]["NotEnoughDiskSpace"];
      };
      requestBody: components["requestBodies"]["File"];
    };
  };
}

export interface components {
  schemas: {
    Error: {
      /** @description Error message */
      message: string;
      /** @description Error code */
      code: number;
    };
  };
  responses: {
    /** The file was uploaded successfully. */
    UploadSuccess: unknown;
    /** Entire file downloaded successfully. */
    DownloadSuccess: {
      content: {
        "application/octet-stream": string;
      };
    };
    /** Directory path error */
    DirectoryPathError: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** Internal server error */
    InternalServerError: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** File not found */
    FileNotFound: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** Invalid user */
    InvalidUser: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
    /** Not enough disk space */
    NotEnoughDiskSpace: {
      content: {
        "application/json": components["schemas"]["Error"];
      };
    };
  };
  parameters: {
    /** @description Path to the file, URL encoded. Can be relative to user's home directory. */
    FilePath: string;
    /** @description User used for setting the owner, or resolving relative paths. */
    User: string;
  };
  requestBodies: {
    File: {
      content: {
        "multipart/form-data": {
          /** Format: binary */
          file?: string;
        };
      };
    };
  };
}

export interface operations {}

export interface external {}