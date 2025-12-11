// @generated
/// Generated client implementations.
pub mod relay_client {
    #![allow(
        unused_variables,
        dead_code,
        missing_docs,
        clippy::wildcard_imports,
        clippy::let_unit_value,
    )]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct RelayClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl RelayClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> RelayClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::Body>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + std::marker::Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + std::marker::Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> RelayClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::Body>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::Body>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::Body>,
            >>::Error: Into<StdError> + std::marker::Send + std::marker::Sync,
        {
            RelayClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn submit_block(
            &mut self,
            request: impl tonic::IntoRequest<super::SubmitBlockRequest>,
        ) -> std::result::Result<
            tonic::Response<super::SubmitBlockResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/SubmitBlock");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "SubmitBlock"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn register_validator(
            &mut self,
            request: impl tonic::IntoRequest<super::RegisterValidatorRequest>,
        ) -> std::result::Result<
            tonic::Response<super::RegisterValidatorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/RegisterValidator");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "RegisterValidator"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_header(
            &mut self,
            request: impl tonic::IntoRequest<super::GetHeaderRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetHeaderResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/GetHeader");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "GetHeader"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_payload(
            &mut self,
            request: impl tonic::IntoRequest<super::GetPayloadRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetPayloadResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/GetPayload");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "GetPayload"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn stream_header(
            &mut self,
            request: impl tonic::IntoRequest<super::StreamHeaderRequest>,
        ) -> std::result::Result<
            tonic::Response<tonic::codec::Streaming<super::StreamHeaderResponse>>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/StreamHeader");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "StreamHeader"));
            self.inner.server_streaming(req, path, codec).await
        }
        pub async fn stream_block(
            &mut self,
            request: impl tonic::IntoRequest<super::StreamBlockRequest>,
        ) -> std::result::Result<
            tonic::Response<tonic::codec::Streaming<super::StreamBlockResponse>>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/StreamBlock");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "StreamBlock"));
            self.inner.server_streaming(req, path, codec).await
        }
        pub async fn forward_block(
            &mut self,
            request: impl tonic::IntoRequest<super::StreamBlockResponse>,
        ) -> std::result::Result<
            tonic::Response<super::SubmitBlockResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/ForwardBlock");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "ForwardBlock"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_validator_registration(
            &mut self,
            request: impl tonic::IntoRequest<super::GetValidatorRegistrationRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetValidatorRegistrationResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/Relay/GetValidatorRegistration",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("Relay", "GetValidatorRegistration"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn pre_fetch_get_payload(
            &mut self,
            request: impl tonic::IntoRequest<super::PreFetchGetPayloadRequest>,
        ) -> std::result::Result<
            tonic::Response<super::PreFetchGetPayloadResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/PreFetchGetPayload");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "PreFetchGetPayload"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn stream_builder(
            &mut self,
            request: impl tonic::IntoRequest<super::StreamBuilderRequest>,
        ) -> std::result::Result<
            tonic::Response<tonic::codec::Streaming<super::StreamBuilderResponse>>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/StreamBuilder");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "StreamBuilder"));
            self.inner.server_streaming(req, path, codec).await
        }
        pub async fn stream_slot_info(
            &mut self,
            request: impl tonic::IntoRequest<super::StreamSlotRequest>,
        ) -> std::result::Result<
            tonic::Response<tonic::codec::Streaming<super::StreamSlotResponse>>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/StreamSlotInfo");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "StreamSlotInfo"));
            self.inner.server_streaming(req, path, codec).await
        }
        pub async fn ping(
            &mut self,
            request: impl tonic::IntoRequest<super::PingRequest>,
        ) -> std::result::Result<tonic::Response<super::PingResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static("/Relay/Ping");
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "Ping"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn send_header_delivered(
            &mut self,
            request: impl tonic::IntoRequest<super::HeaderDeliveredRequest>,
        ) -> std::result::Result<
            tonic::Response<super::HeaderDeliveredResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/Relay/SendHeaderDelivered",
            );
            let mut req = request.into_request();
            req.extensions_mut().insert(GrpcMethod::new("Relay", "SendHeaderDelivered"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn adjust_latest_block_payload(
            &mut self,
            request: impl tonic::IntoRequest<super::AdjustLatestBlockPayloadRequest>,
        ) -> std::result::Result<
            tonic::Response<super::AdjustLatestBlockPayloadResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic_prost::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/Relay/AdjustLatestBlockPayload",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("Relay", "AdjustLatestBlockPayload"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod relay_server {
    #![allow(
        unused_variables,
        dead_code,
        missing_docs,
        clippy::wildcard_imports,
        clippy::let_unit_value,
    )]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with RelayServer.
    #[async_trait]
    pub trait Relay: std::marker::Send + std::marker::Sync + 'static {
        async fn submit_block(
            &self,
            request: tonic::Request<super::SubmitBlockRequest>,
        ) -> std::result::Result<
            tonic::Response<super::SubmitBlockResponse>,
            tonic::Status,
        >;
        async fn register_validator(
            &self,
            request: tonic::Request<super::RegisterValidatorRequest>,
        ) -> std::result::Result<
            tonic::Response<super::RegisterValidatorResponse>,
            tonic::Status,
        >;
        async fn get_header(
            &self,
            request: tonic::Request<super::GetHeaderRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetHeaderResponse>,
            tonic::Status,
        >;
        async fn get_payload(
            &self,
            request: tonic::Request<super::GetPayloadRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetPayloadResponse>,
            tonic::Status,
        >;
        /// Server streaming response type for the StreamHeader method.
        type StreamHeaderStream: tonic::codegen::tokio_stream::Stream<
                Item = std::result::Result<super::StreamHeaderResponse, tonic::Status>,
            >
            + std::marker::Send
            + 'static;
        async fn stream_header(
            &self,
            request: tonic::Request<super::StreamHeaderRequest>,
        ) -> std::result::Result<
            tonic::Response<Self::StreamHeaderStream>,
            tonic::Status,
        >;
        /// Server streaming response type for the StreamBlock method.
        type StreamBlockStream: tonic::codegen::tokio_stream::Stream<
                Item = std::result::Result<super::StreamBlockResponse, tonic::Status>,
            >
            + std::marker::Send
            + 'static;
        async fn stream_block(
            &self,
            request: tonic::Request<super::StreamBlockRequest>,
        ) -> std::result::Result<
            tonic::Response<Self::StreamBlockStream>,
            tonic::Status,
        >;
        async fn forward_block(
            &self,
            request: tonic::Request<super::StreamBlockResponse>,
        ) -> std::result::Result<
            tonic::Response<super::SubmitBlockResponse>,
            tonic::Status,
        >;
        async fn get_validator_registration(
            &self,
            request: tonic::Request<super::GetValidatorRegistrationRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetValidatorRegistrationResponse>,
            tonic::Status,
        >;
        async fn pre_fetch_get_payload(
            &self,
            request: tonic::Request<super::PreFetchGetPayloadRequest>,
        ) -> std::result::Result<
            tonic::Response<super::PreFetchGetPayloadResponse>,
            tonic::Status,
        >;
        /// Server streaming response type for the StreamBuilder method.
        type StreamBuilderStream: tonic::codegen::tokio_stream::Stream<
                Item = std::result::Result<super::StreamBuilderResponse, tonic::Status>,
            >
            + std::marker::Send
            + 'static;
        async fn stream_builder(
            &self,
            request: tonic::Request<super::StreamBuilderRequest>,
        ) -> std::result::Result<
            tonic::Response<Self::StreamBuilderStream>,
            tonic::Status,
        >;
        /// Server streaming response type for the StreamSlotInfo method.
        type StreamSlotInfoStream: tonic::codegen::tokio_stream::Stream<
                Item = std::result::Result<super::StreamSlotResponse, tonic::Status>,
            >
            + std::marker::Send
            + 'static;
        async fn stream_slot_info(
            &self,
            request: tonic::Request<super::StreamSlotRequest>,
        ) -> std::result::Result<
            tonic::Response<Self::StreamSlotInfoStream>,
            tonic::Status,
        >;
        async fn ping(
            &self,
            request: tonic::Request<super::PingRequest>,
        ) -> std::result::Result<tonic::Response<super::PingResponse>, tonic::Status>;
        async fn send_header_delivered(
            &self,
            request: tonic::Request<super::HeaderDeliveredRequest>,
        ) -> std::result::Result<
            tonic::Response<super::HeaderDeliveredResponse>,
            tonic::Status,
        >;
        async fn adjust_latest_block_payload(
            &self,
            request: tonic::Request<super::AdjustLatestBlockPayloadRequest>,
        ) -> std::result::Result<
            tonic::Response<super::AdjustLatestBlockPayloadResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct RelayServer<T> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T> RelayServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for RelayServer<T>
    where
        T: Relay,
        B: Body + std::marker::Send + 'static,
        B::Error: Into<StdError> + std::marker::Send + 'static,
    {
        type Response = http::Response<tonic::body::Body>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/Relay/SubmitBlock" => {
                    #[allow(non_camel_case_types)]
                    struct SubmitBlockSvc<T: Relay>(pub Arc<T>);
                    impl<T: Relay> tonic::server::UnaryService<super::SubmitBlockRequest>
                    for SubmitBlockSvc<T> {
                        type Response = super::SubmitBlockResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::SubmitBlockRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::submit_block(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = SubmitBlockSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/RegisterValidator" => {
                    #[allow(non_camel_case_types)]
                    struct RegisterValidatorSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::UnaryService<super::RegisterValidatorRequest>
                    for RegisterValidatorSvc<T> {
                        type Response = super::RegisterValidatorResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::RegisterValidatorRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::register_validator(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RegisterValidatorSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/GetHeader" => {
                    #[allow(non_camel_case_types)]
                    struct GetHeaderSvc<T: Relay>(pub Arc<T>);
                    impl<T: Relay> tonic::server::UnaryService<super::GetHeaderRequest>
                    for GetHeaderSvc<T> {
                        type Response = super::GetHeaderResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetHeaderRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::get_header(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetHeaderSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/GetPayload" => {
                    #[allow(non_camel_case_types)]
                    struct GetPayloadSvc<T: Relay>(pub Arc<T>);
                    impl<T: Relay> tonic::server::UnaryService<super::GetPayloadRequest>
                    for GetPayloadSvc<T> {
                        type Response = super::GetPayloadResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetPayloadRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::get_payload(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetPayloadSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/StreamHeader" => {
                    #[allow(non_camel_case_types)]
                    struct StreamHeaderSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::ServerStreamingService<super::StreamHeaderRequest>
                    for StreamHeaderSvc<T> {
                        type Response = super::StreamHeaderResponse;
                        type ResponseStream = T::StreamHeaderStream;
                        type Future = BoxFuture<
                            tonic::Response<Self::ResponseStream>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::StreamHeaderRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::stream_header(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = StreamHeaderSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.server_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/StreamBlock" => {
                    #[allow(non_camel_case_types)]
                    struct StreamBlockSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::ServerStreamingService<super::StreamBlockRequest>
                    for StreamBlockSvc<T> {
                        type Response = super::StreamBlockResponse;
                        type ResponseStream = T::StreamBlockStream;
                        type Future = BoxFuture<
                            tonic::Response<Self::ResponseStream>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::StreamBlockRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::stream_block(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = StreamBlockSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.server_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/ForwardBlock" => {
                    #[allow(non_camel_case_types)]
                    struct ForwardBlockSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::UnaryService<super::StreamBlockResponse>
                    for ForwardBlockSvc<T> {
                        type Response = super::SubmitBlockResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::StreamBlockResponse>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::forward_block(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ForwardBlockSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/GetValidatorRegistration" => {
                    #[allow(non_camel_case_types)]
                    struct GetValidatorRegistrationSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::UnaryService<super::GetValidatorRegistrationRequest>
                    for GetValidatorRegistrationSvc<T> {
                        type Response = super::GetValidatorRegistrationResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::GetValidatorRegistrationRequest,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::get_validator_registration(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetValidatorRegistrationSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/PreFetchGetPayload" => {
                    #[allow(non_camel_case_types)]
                    struct PreFetchGetPayloadSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::UnaryService<super::PreFetchGetPayloadRequest>
                    for PreFetchGetPayloadSvc<T> {
                        type Response = super::PreFetchGetPayloadResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::PreFetchGetPayloadRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::pre_fetch_get_payload(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = PreFetchGetPayloadSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/StreamBuilder" => {
                    #[allow(non_camel_case_types)]
                    struct StreamBuilderSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::ServerStreamingService<super::StreamBuilderRequest>
                    for StreamBuilderSvc<T> {
                        type Response = super::StreamBuilderResponse;
                        type ResponseStream = T::StreamBuilderStream;
                        type Future = BoxFuture<
                            tonic::Response<Self::ResponseStream>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::StreamBuilderRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::stream_builder(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = StreamBuilderSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.server_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/StreamSlotInfo" => {
                    #[allow(non_camel_case_types)]
                    struct StreamSlotInfoSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::ServerStreamingService<super::StreamSlotRequest>
                    for StreamSlotInfoSvc<T> {
                        type Response = super::StreamSlotResponse;
                        type ResponseStream = T::StreamSlotInfoStream;
                        type Future = BoxFuture<
                            tonic::Response<Self::ResponseStream>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::StreamSlotRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::stream_slot_info(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = StreamSlotInfoSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.server_streaming(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/Ping" => {
                    #[allow(non_camel_case_types)]
                    struct PingSvc<T: Relay>(pub Arc<T>);
                    impl<T: Relay> tonic::server::UnaryService<super::PingRequest>
                    for PingSvc<T> {
                        type Response = super::PingResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::PingRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::ping(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = PingSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/SendHeaderDelivered" => {
                    #[allow(non_camel_case_types)]
                    struct SendHeaderDeliveredSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::UnaryService<super::HeaderDeliveredRequest>
                    for SendHeaderDeliveredSvc<T> {
                        type Response = super::HeaderDeliveredResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::HeaderDeliveredRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::send_header_delivered(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = SendHeaderDeliveredSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/Relay/AdjustLatestBlockPayload" => {
                    #[allow(non_camel_case_types)]
                    struct AdjustLatestBlockPayloadSvc<T: Relay>(pub Arc<T>);
                    impl<
                        T: Relay,
                    > tonic::server::UnaryService<super::AdjustLatestBlockPayloadRequest>
                    for AdjustLatestBlockPayloadSvc<T> {
                        type Response = super::AdjustLatestBlockPayloadResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::AdjustLatestBlockPayloadRequest,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Relay>::adjust_latest_block_payload(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = AdjustLatestBlockPayloadSvc(inner);
                        let codec = tonic_prost::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        let mut response = http::Response::new(
                            tonic::body::Body::default(),
                        );
                        let headers = response.headers_mut();
                        headers
                            .insert(
                                tonic::Status::GRPC_STATUS,
                                (tonic::Code::Unimplemented as i32).into(),
                            );
                        headers
                            .insert(
                                http::header::CONTENT_TYPE,
                                tonic::metadata::GRPC_CONTENT_TYPE,
                            );
                        Ok(response)
                    })
                }
            }
        }
    }
    impl<T> Clone for RelayServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    /// Generated gRPC service name
    pub const SERVICE_NAME: &str = "Relay";
    impl<T> tonic::server::NamedService for RelayServer<T> {
        const NAME: &'static str = SERVICE_NAME;
    }
}
