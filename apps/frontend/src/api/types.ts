import {apiContract} from "@taskForge/openapi/contracts"
import type {ServerInferRequest} from "@ts-rest/core";

export type TRequests = ServerInferRequest<typeof apiContract>;